package com

type ResModel struct {
	R       interface{} `json:"r"`
	Success string      `json:"success"`
	Code    int         `json:"code"`
	Info    string      `json:"info"`
}

type PageResModel struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
