#
# Node2D
#     KinematicBody2D
#         Sprite
#         CollisionShape2D
#

extends KinematicBody2D

const speed : int = 200

var velocity : Vector2
var need_rotate : float # radians

onready var target : Vector2 = self.position
	
func _input(event):
	if event is InputEventMouseButton:
		if event.button_index == BUTTON_LEFT and event.pressed:
			target = get_global_mouse_position()

func _physics_process(_delta):
	velocity = position.direction_to(target)
	
	if position.distance_to(target) > 5:
		need_rotate = self.get_angle_to(target)
		
		if need_rotate > 0.1:
			self.rotation += 0.05

		elif need_rotate < -0.1:
			self.rotation -= 0.05
            
		else:
			velocity = move_and_slide(velocity * speed)
