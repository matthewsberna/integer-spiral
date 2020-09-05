package main

import (
    "flag"
    "fmt"
)

type Coordinate struct {
    X, Y int
}

type LayerProperties struct {
    // Defines the boundaries of the four legs of our layer.
    leg0Start   Coordinate
    leg0End     Coordinate
    leg1Start   Coordinate
    leg1End     Coordinate
    leg2Start   Coordinate
    leg2End     Coordinate
    leg3Start   Coordinate
    leg3End     Coordinate
}

// Globals to hold stuff, so I don't have to pass that s*** around constantly.  The functions have enough arguments already.
var N                       int
var NumberOfLayers          int
var LayerPropertiesCache    map[int]LayerProperties


func main() {
    sizePtr := flag.Int("size", 0, "The size of the matrix")
    flag.Parse()
    
    // Initialize the global N.
    N = *sizePtr
    // You're not even trying.
    if !(N > 0) {
        fmt.Println("You must provide a valid size.  Use -size=<someRealNumber>.")
        return
    }
    fmt.Printf("N = %v\n", N)
    
    // Initialize the globals.
    if N %2 != 0 {
        NumberOfLayers = N/2 + 1
    } else {
        NumberOfLayers = N/2
    }
    
    LayerPropertiesCache = make(map[int]LayerProperties)

    // Construct our matrix.
    matrix := make([][]int, N)
    for i := range matrix {
        matrix[i] = make([]int, N)
    }
    
    // Populate the matrix with values.
    for x := 0; x < N; x++ {
        for y := 0; y < N; y++ {
            layerNumber := CalculateLayerNumber(x, y)
            legNumber := CalculateLegNumber(x, y, layerNumber)
            legOffset := CalculateLegOffset(x, y, layerNumber, legNumber)

            elementValue := CalculateElementValue(layerNumber, legNumber, legOffset)
            matrix[x][y] = elementValue
        }
    }
    
    // Print it.
    for x := 0; x < N; x++ {
        for y := 0; y < N; y++ {
            fmt.Printf("%-3d", matrix[x][y])
        }
        fmt.Println()
    }
}

func CalculateElementValue(layerNumber int, legNumber int, legOffset int) int {
    elementValue := 0
    
    // Add elements of completed layers.
    for m := 0; m < layerNumber; m++ {
        elementValue += ( 4*(N - 2*m - 1) )
    }

    // Add elements of completed legs.
    for i := 0; i < legNumber; i++ {
        elementValue += (N - 2*layerNumber - 1)
    }

    // Finally, add the offset.
    elementValue += legOffset

    return elementValue
}

func CalculateLayerNumber(X int, Y int) int {
    var x, y int
    
    if X > NumberOfLayers - 1 {
        x = N - X - 1
    } else {
        x = X
    }
    if Y > NumberOfLayers - 1 {
        y = N - Y - 1
    } else {
        y = Y
    }

    layerNum := Min(x, y)
    
    return layerNum
}

func CalculateLegNumber(X int, Y int, LayerNum int) int {
    layerProps := GetLayerProperties(LayerNum)

    var legNumber int

    if (X >= layerProps.leg0Start.X && X <= layerProps.leg0End.X && Y == layerProps.leg0Start.Y) {
        legNumber = 0
    } else if (X == layerProps.leg1Start.X && Y >= layerProps.leg1Start.Y && Y <= layerProps.leg1End.Y) {
        legNumber = 1
    } else if (X >= layerProps.leg2Start.X && X <= layerProps.leg2End.X && Y == layerProps.leg2Start.Y) {
        legNumber = 2
    } else if (X == layerProps.leg3Start.X && Y >= layerProps.leg3Start.Y && Y <= layerProps.leg3End.Y) {
        legNumber = 3
    } else {
        // Well where the hell are we, then??  Probably wanna throw an exception at this point, huh?
    }

    return legNumber
}

func CalculateLegOffset(X int, Y int, LayerNum int, LegNum int) int {
    // Special case for the innermost layer.
    if LayerNum == (NumberOfLayers - 1) {
        return 0
    }
    
    layerProps := GetLayerProperties(LayerNum)

    switch LegNum {
        case 0:
            return X - layerProps.leg0Start.X
        case 1:
            return Y - layerProps.leg1Start.Y
        case 2:
            return layerProps.leg2End.X - X
        case 3:
            return layerProps.leg3End.Y - Y
        default:
            // Danger, Will Robinson!
            // Throw an exception if you have 5 sides in your square.
    }
    return -1
}


// Helper functions

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func GetLayerProperties(LayerNum int) LayerProperties {
    layerProps, found := LayerPropertiesCache[LayerNum]
    if !found {
        layerProps := LayerProperties {
            leg0Start : Coordinate{X: LayerNum, Y: LayerNum},
            leg0End   : Coordinate{X: N - LayerNum - 2, Y: LayerNum},
            leg1Start : Coordinate{X: N - LayerNum - 1, Y: LayerNum},
            leg1End   : Coordinate{X: N - LayerNum - 1, Y: N - LayerNum - 2},
            leg2Start : Coordinate{X: LayerNum + 1, Y: N - LayerNum - 1},
            leg2End   : Coordinate{X: N - LayerNum - 1, Y: N - LayerNum - 1},
            leg3Start : Coordinate{X: LayerNum, Y: LayerNum + 1},
            leg3End   : Coordinate{X: LayerNum, Y: N - LayerNum - 1},
        }
        LayerPropertiesCache[LayerNum] = layerProps
    }
    
    return layerProps
}

func (lp LayerProperties) String() string {
    return fmt.Sprintf("%v-%v\r\n%v-%v\r\n%v-%v\r\n%v-%v\r\n",
        lp.leg0Start, lp.leg0End, lp.leg1Start, lp.leg1End, lp.leg2Start, lp.leg2End, lp.leg3Start, lp.leg3End)
}