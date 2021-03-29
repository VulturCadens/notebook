import bpy

#
# bpy.ops.object.simple_operator_example()
#

class OperatorExample(bpy.types.Operator):
    bl_idname = "object.simple_operator_example"
    bl_label = "Operator Example"

    def execute(self, context):
        print(self.bl_label)
        return {'FINISHED'}


def register():
    bpy.utils.register_class(OperatorExample)

def unregister():
    bpy.utils.unregister_class(OperatorExample)


if __name__ == "__main__":
    register() 

