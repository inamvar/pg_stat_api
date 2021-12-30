/*
 Author: Iman Namvar
 Email: iman.namvar@gmail.com
 */

package main

import "pg_stats/server"

func main() {
  apiServer := &server.ApiServer{}
  apiServer.Start()
  defer  apiServer.ShutDown()
}
