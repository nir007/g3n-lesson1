package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"time"
)

func main()  {
	a := app.App()
	scene := core.NewNode()
	gui.Manager().Set(scene)
	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 4)
	scene.Add(cam)
	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	txSpace, err := texture.NewTexture2DFromImage("./textures/box.jpg")
	if err != nil {
		panic(err)
	}

	g := geometry.NewBox(2, 2, 2)
	mat := material.NewStandard(math32.NewColor("White"))
	mat.AddTexture(txSpace)

	cube := graphic.NewMesh(g, mat)
	cube.RotateX(90)
	cube.RotateY(40)
	cube.RotateZ(-80)

	scene.Add(cube)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}