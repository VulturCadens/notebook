extends Node2D

var mousepos : Vector2

func _ready():
	pass
	
func _input(event):
	if event is InputEventMouseButton:
		if event.button_index == MOUSE_BUTTON_LEFT and event.pressed:
			mousepos = get_global_mouse_position()
			print(mousepos)
			
			Input.warp_mouse(Vector2(mousepos.x - 10, mousepos.y))

func _process(_delta):
	pass
