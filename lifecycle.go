package main

type Lifecycle interface {
	OnInit()
	OnUpdate()
	OnDraw()
}
