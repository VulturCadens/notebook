#
# The first way is to use a PathFollow2D node.
#
# Path2D (Contains a Curve2D path.)
#     PathFollow2D (Offset property is the distance along the path.)
#         KinematicBody2D  
#             Sprite
#             CollisionShape2D
#

extends KinematicBody2D

const speed = 100

onready var parent : PathFollow2D = get_parent()

func _physics_process(delta):
	parent.set_offset(parent.get_offset() + speed * delta)

#
# The second way is to use the path as the target.
#
# Path2D (Contains a Curve2D path.)
#     KinematicBody2D  
#         Sprite
#         CollisionShape2D
#

extends KinematicBody2D

const speed = 100

var index : int
var max_index : int
var points : PoolVector2Array
var target : Vector2
var _velocity : Vector2

func _ready():
	points = get_parent().curve.get_baked_points()
	index = 0
	max_index = points.size()
	target = points[index]
		
func _physics_process(_delta):
	if self.position.distance_to(target) < 1:
		index = wrapi(index + 1, 0, max_index)
		target = points[index]
	
	var direction = (target - self.position).normalized()
	
	_velocity = move_and_slide(direction * speed)
