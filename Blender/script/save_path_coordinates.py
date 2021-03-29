import bpy
import math

from bpy.types import Operator
from bpy_extras.io_utils import ExportHelper

from dataclasses import dataclass

# Property Definitions (bpy.props)
# 
# The result of these functions is used to assign properties
# to classes registered with Blender and can’t be used directly.

from bpy.props import StringProperty, IntProperty


bl_info = {
    "name": "Path Coordinates",
    "author": "Vultur Cadens",
    "version": (0, 1),
    "blender": (2, 92, 0),
    "location": "View3D > Object",
    "category": "Object",
}


@dataclass
class Point:

    x: int
    y: int


class SavePathCoordinates(Operator, ExportHelper):

    bl_label = "Save Path Coordinates"
    bl_idname = "object.save_path_coordinates"

    # ExportHelper mixin class uses these.
    
    resolution = IntProperty(
        name="Resolution",
        default=10,
        min=5,
        max=100
    )
    
    filename_ext = ".csv" 

    filter_glob = StringProperty(
        default="*.csv",
        options={"HIDDEN"},
        maxlen=255
    )

    def execute(self, context):
        bpy.ops.object.mode_set(mode="OBJECT")

        originalCurve = bpy.context.active_object

        if not originalCurve:
            print("There isn't an active object.")
            return {"FINISHED"}

        if originalCurve.type != "CURVE":
            print("The active object isn't a curve.")
            return {"FINISHED"}

        original_name = originalCurve.name

        temporaryCurve = originalCurve.copy()
        temporaryCurve.data = originalCurve.data.copy()
        temporaryCurve.name = "CurveTemporaryClone"

        bpy.context.scene.collection.objects.link(temporaryCurve)

        bpy.ops.object.select_all(action="DESELECT")

        bpy.data.objects["CurveTemporaryClone"].select_set(True)
        bpy.context.view_layer.objects.active = temporaryCurve

        bpy.context.object.data.resolution_u = self.resolution
        bpy.ops.object.convert(target="MESH")

        points = []

        for vertex in temporaryCurve.data.vertices:
            points.append(Point(
                vertex.co[0],
                vertex.co[1]
            ))

        points = self.normal(points)

        self.write(context, self.filepath, points)

        bpy.ops.object.delete()

        bpy.data.objects[original_name].select_set(True)
        bpy.context.view_layer.objects.active = originalCurve

        return {"FINISHED"}

    def write(self, context, filepath, points):
        with open(filepath, "w+", encoding="utf8") as file:
            for point in points:
                file.write(str(point.x) + "," + str(point.y) + ",\n")

    def normal(self, points):
        min_x = max_x = points[0].x
        min_y = max_y = points[0].y

        for point in points:
            if point.x < min_x:
                min_x = point.x
            elif point.x > max_x:
                max_x = point.x

            if point.y < min_y:
                min_y = point.y
            elif point.y > max_y:
                max_y = point.y

        normal_points = []

        for point in points:
            x = self.map(point.x, min_x, max_x)
            y = self.map(point.y, min_y, max_y)

            normal_points.append(Point(x, y))

        return normal_points

    def map(self, value, min, max):
        return (value - min) / (max - min)

#
# Registration
#


def register():
    bpy.utils.register_class(SavePathCoordinates)
    bpy.types.VIEW3D_MT_object.append(add_button)


def unregister():
    bpy.utils.unregister_class(SavePathCoordinates)
    bpy.types.VIEW3D_MT_object.remove(add_button)


def add_button(self, context):
    self.layout.operator(
        SavePathCoordinates.bl_idname,
        text="Save Path Coordinates",
        icon="PLUGIN"
    )


if __name__ == "__main__":
    register()
