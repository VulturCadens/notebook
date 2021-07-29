import bpy

#
# bpy.ops.object.simple_operator_example()
#


class OperatorExample(bpy.types.Operator):
    bl_idname = "object.simple_operator_example"
    bl_label = "Operator Example"

    def execute(self, context):
        print(self.bl_label)

        if not context.selected_objects:
            print("The collection is empty")
            return {"FINISHED"}

        for obj in context.selected_objects:
            print(obj.name)

        return {"FINISHED"}


def register():
    bpy.utils.register_class(OperatorExample)


def unregister():
    bpy.utils.unregister_class(OperatorExample)


if __name__ == "__main__":
    register()
