package Info

import (
	"net"
)

type UserConnection struct {
	Conn net.Conn
	logged bool
	UserInfo User
}

func NewUserConnection(conn net.Conn, userinfo User) UserConnection {
	return UserConnection{Conn : conn, logged: false, UserInfo:userinfo}
}

func (coo *UserConnection) HasLogged() bool {
	return coo.logged
}

func (coo *UserConnection) LoggedIn() {
	coo.logged = true
}
