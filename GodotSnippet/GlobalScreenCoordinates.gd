#
# Get the global coordinates of the top of the screen.
#

func get_camera_rect(camera : Camera2D) -> Rect2:
	var cameraPos : Vector2 = camera.get_camera_screen_center() # Relative to the world origin.
	var viewportSize : Vector2 = get_viewport_rect().size * camera.zoom

	var position : Vector2 = Vector2(cameraPos.x - (viewportSize.x / 2), cameraPos.y - (viewportSize.y / 2))
	var size : Vector2 = viewportSize

	return Rect2(position, size)
    