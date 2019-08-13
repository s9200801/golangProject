package http

type Handler interface{
	ServerHttp(w ResponseWriter, r *Requset)
}