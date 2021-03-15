#
# bullet.tscn
#
# Area2D
#     Sprite
#     CollisionShape2D
#     VisibilityNotifier2D
#

extends Area2D

const speed : int = 500

func _physics_process(delta):
	self.position -= self.transform.y * speed * delta

#Signal VisibilityNotifier::screen_exited() ->
func _on_VisibilityNotifier2D_screen_exited():
	self.queue_free()

# Signal Area2D::body_entered(body : Node) ->
func _on_Area2D_body_entered(_body):
	self.queue_free()

#
# const bullet = preload("res://bullet.tscn") 
#
# var b = bullet.instance()
# get_node("/root").add_child(b)
# b.transform = $Position2D.global_transform
#