package gosupervisor

import (
	"github.com/kolo/xmlrpc"
)

// getAPIVersion()
// Return the version of the RPC API used by supervisord
//
// @return string version version id
func (rpc *SupervisorRpc) GetAPIVersion() (string, error) {
	client, err := xmlrpc.NewClient(rpc.Url, rpc.Transport)
	ret := ""
	err = client.Call("supervisor.getAPIVersion", nil, &ret)
	defer client.Close()
	return ret, err
}

//@return boolean result always returns True unless error
func (rpc *SupervisorRpc) Shutdown() (bool, error) {
	client, err := xmlrpc.NewClient(rpc.Url, rpc.Transport)
	ret := false
	err = client.Call("supervisor.shutdown", nil, &ret)
	defer client.Close()
	return ret, err
}

//@return boolean result always returns True unless error
func (rpc *SupervisorRpc) Restart() (bool, error) {
	client, err := xmlrpc.NewClient(rpc.Url, rpc.Transport)
	ret := false
	err = client.Call("supervisor.restart", nil, &ret)
	defer client.Close()
	return ret, err
}