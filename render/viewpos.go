package render

import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"log"
	"math"
)

var (
	Viewpos                [3]float32
	ViewPhi, ViewTheta     float64
	mousePrevX, mousePrevY int
	mouseButton            [5]int
)

var Width, Height int

const PI = math.Pi

func InitViewport() {
	gl.Viewport(0, 0, gl.Sizei(Width), gl.Sizei(Height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	x := gl.Double(float64(Height) / float64(Width))
	gl.Frustum(-1./2., 1./2., -x/2, x/2, 10, 100)
}

// Set the GL modelview matrix to match view position and orientation.
func UpdateViewpos() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(gl.Float(Viewpos[0]), gl.Float(Viewpos[1]), gl.Float(Viewpos[2]))
	gl.Rotatef(gl.Float(ViewTheta*(180/PI)), 1, 0, 0)
	gl.Rotatef(gl.Float(ViewPhi*(180/PI)), 0, 0, 1)
}

const (
	deltaMove = 0.1
	deltaLook = 0.01
)

// Sets up input handlers
func InitInputHandlers() {
	glfw.SetKeyCallback(func(key, state int) {
		if state == 1 {
			switch key {
			case Up:
				Viewpos[2] += deltaMove
			case Down:
				Viewpos[2] -= deltaMove
			case Left:
				Viewpos[1] += deltaMove
			case Right:
				Viewpos[1] -= deltaMove
			case Space:
				Viewpos[0] += deltaMove
			case Alt:
				Viewpos[0] -= deltaMove
			default:
				log.Println("unused key:", key)
			}
		}
	})

	glfw.SetMousePosCallback(func(x, y int) {

		dx, dy := x-mousePrevX, y-mousePrevY
		mousePrevX, mousePrevY = x, y
		if mouseButton[0] == 0 {
			return
		}

		ViewPhi += deltaLook * float64(dx)
		ViewTheta += deltaLook * float64(dy)

		// limit viewing angles
		if ViewPhi < -PI {
			ViewPhi += 2 * PI
		}
		if ViewPhi > PI {
			ViewPhi -= 2 * PI
		}
		if ViewTheta > PI {
			ViewTheta = -PI
		}
		if ViewTheta < -PI {
			ViewTheta = PI
		}
		log.Println("phi, theta:", ViewPhi, ViewTheta)
	})

	glfw.SetMouseButtonCallback(func(button, state int) {
		log.Println("mousebutton:", button, state)
		mouseButton[button] = state
	})

	glfw.SetMouseWheelCallback(func(delta int) {
		log.Println("mousewheel:", delta)
		glfw.SetMouseWheel(0)
	})

	glfw.SetWindowSizeCallback(func(w, h int) {
		Width, Height = w, h
		InitViewport()
	})
}

const (
	Up    = 283
	Down  = 284
	Left  = 285
	Right = 286
	Space = 32
	Alt   = 291
	Esc   = 257
)
