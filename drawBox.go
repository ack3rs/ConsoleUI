package ConsoleUI

func DrawBox(x int, y int, w int, h int, style string) {
    
    myStyle := getStyle(style)
    
    MoveTo(x, y)
    
    // maxLength := 1
    
    // Draw a Box Starting at X, Y with the MaxLength
    // Out(myStyle.TopLeft)
    // Out(myStyle.HeaderText)
    // Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-(len(myStyle.HeaderText))))
    // Out(myStyle.TopRight)
    
    // for i := 1; i <= myStyle.SpaceHeader; i++ {
    //    MoveTo(x, y+i)
    //    Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
    // }
    
}
