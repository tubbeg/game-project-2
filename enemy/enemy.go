components {
  id: "enemy"
  component: "/enemy/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"banana\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/enemy/enemy.atlas\"\n"
  "}\n"
  ""
}
