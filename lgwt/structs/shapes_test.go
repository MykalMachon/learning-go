package structs

import "testing"

func TestPermiter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimiter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("area of a rectangle", func(t *testing.T) {
	// 	rect := Rectangle{12.0, 6.0}
	// 	checkArea(t, rect, 72.0)
	// })

	// t.Run("area of a circle", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	// these are called "table tests"
	// we use an "anonymous struct" to define the shape of the slice
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.want {
				t.Errorf("%#v got %g but want %g", testCase.shape, got, testCase.want)
			}
		})
	}
}
