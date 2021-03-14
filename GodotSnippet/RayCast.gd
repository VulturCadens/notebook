#
# RayCast2D
#
# Use global coordinates.
#

extends RayCast2D

var target : Vector2 
var state : Physics2DDirectSpaceState
var result : Dictionary

func _ready():
	target = Vector2(460, 500)

func _physics_process(_delta):
	state  = get_world_2d().direct_space_state
	result = state.intersect_ray(
		self.global_position,      # from
		target,                    # to
		[self],                    # collision exceptions
		self.collision_mask        # collision mask
	)

	if !result.empty():
		print(result)

#
# The result dictionary when a collision occurs.
#
#   position: Vector2     # point in world space for collision
#   normal: Vector2       # normal in world space for collision
#   collider: Object      # Object collided or null (if unassociated)
#   collider_id: ObjectID # Object it collided against
#   rid: RID              # RID it collided against
#   shape: int            # shape index of collider
#   metadata: Variant()   # metadata of collider
#
# https://docs.godotengine.org/en/stable/tutorials/physics/ray-casting.html
#