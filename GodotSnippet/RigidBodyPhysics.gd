#
# RigidBody2D
#     Sprite
#     CollisionShape2D
#

extends RigidBody2D

const offset = Vector2(0, 0) # offset from the object’s origin

onready var force = Vector2(0.5, -4).normalized() *  100 # direction and magnitude

var is_kick : bool = false

func _ready():
	self.mode = MODE_RIGID # default

func _input(event):
	if event is InputEventKey and event.pressed:
		if event.scancode == KEY_K:
			is_kick = true
			
func _integrate_forces(_state):
	if is_kick:
		is_kick = false
		
		self.apply_torque_impulse(2000)
		
		self.apply_impulse(offset, force)
