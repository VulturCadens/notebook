import bpy
import bmesh
import mathutils

def main():
    if bpy.context.scene.objects.get("BMesh_Object"):
        for object in bpy.context.scene.objects:
            if object.name == "BMesh_Object":
                bpy.data.objects.remove(object, do_unlink = True)
                print("BMesh_Object has been deleted.")
        
        for mesh in bpy.data.meshes:
            if mesh.users == 0:
                bpy.data.meshes.remove(mesh)
                print("BMesh_Mesh has been deleted.")


    # Create an empty BMesh and add a cube.
    bm = bmesh.new() 
    bmesh.ops.create_cube(bm, size = 1)

    
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


    # Create a mesh and write BMesh to the mesh. Free BMesh and prevent further access.
    cube_mesh = bpy.data.meshes.new("BMesh_Mesh")
    bm.to_mesh(cube_mesh)
    bm.free()

    cube_object = bpy.data.objects.new("BMesh_Object", cube_mesh)
    bpy.context.collection.objects.link(cube_object)


if __name__ == "__main__":
    main()
    