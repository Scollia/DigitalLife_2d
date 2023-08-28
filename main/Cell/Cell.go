components {
  id: "Cell"
  component: "/main/Cell/Cell.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite_Cell_(biology)"
  type: "sprite"
  data: "tile_set: \"/main/Cell/Cell.atlas\"\n"
  "default_animation: \"Cell_(biology)\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_MULT\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite_Nucleolis_(biology)"
  type: "sprite"
  data: "tile_set: \"/main/Cell/Cell.atlas\"\n"
  "default_animation: \"Nucleolis_(biology)\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite_anim"
  type: "sprite"
  data: "tile_set: \"/builtins/graphics/particle_blob.tilesource\"\n"
  "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 40.0
    y: 40.0
    z: 1.0
  }
}
