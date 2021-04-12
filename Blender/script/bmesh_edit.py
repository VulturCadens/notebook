import bpy
import bmesh
import mathutils

def main():
    if not bpy.context.scene.objects.get("Cube"):
        print("The 'Cube' object doesn't exist.")
        return

    cube_mesh = bpy.data.objects["Cube"].data

    # Create an empty BMesh and fill it in from a mesh.
    bm = bmesh.new() 
    bm.from_mesh(cube_mesh)
    bpy.ops.object.mode_set(mode="EDIT")

    
    bm.faces.ensure_lookup_table()    
    
    # Extrude operator (does not transform).
    returns = bmesh.ops.extrude_face_region(
        bm,
        geom = [bm.faces[1]]
    )
    
    vertices = [v for v in returns["geom"] if isinstance(v, bmesh.types.BMVert)]

    bmesh.ops.translate(
        bm,
        vec = mathutils.Vector((0, 1.2, 0)),
        verts = vertices
    )


    # Write the bmesh back to the mesh. Free and prevent further access.
    bpy.ops.object.mode_set(mode="OBJECT")
    bm.to_mesh(cube_mesh)
    bm.free()

if __name__ == "__main__":    
    main()
    
    