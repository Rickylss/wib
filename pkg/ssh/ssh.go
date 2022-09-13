package ssh

package node

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (n *Node) connect(cipherList []string) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	if n.Key == "" {
		auth = append(auth, ssh.Password(n.Password))
	} else {
		pemBytes, err := ioutil.ReadFile(n.Key)
		if err != nil {
			n.loge.Error(err)
			return nil, err
		}

		var signer ssh.Signer
		if n.Password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(n.Password))
		}
		if err != nil {
			n.loge.Error(err)
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    n.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", n.Hostname, n.SshPort)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		n.loge.Error(err)
		return nil, err
	}

	return client, nil
}

func (n *Node) GetSession(cipherList []string) (*ssh.Session, error) {
	client, err := n.connect(cipherList)
	if err != nil {
		n.loge.Error(err)
		return nil, err
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		n.loge.Error(err)
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		n.loge.Error(err)
		return nil, err
	}

	return session, nil
}

func (n *Node) GetSftpClient(cipherList []string) (*sftp.Client, error) {
	client, err := n.connect(cipherList)
	if err != nil {
		n.loge.Error(err)
		return nil, err
	}

	// create sftp client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		n.loge.Error(err)
		return nil, err
	}

	return sftpClient, nil
}
