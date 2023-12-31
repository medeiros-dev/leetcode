func canPlaceFlowers(flowerbed []int, n int) bool {

    var avb_plots int 

    if len(flowerbed) == 1{
        avb_plots = 1 - flowerbed[0] 
        goto finish
    } 

    for i:= 0; i < len(flowerbed); i++{
        
        if i == 0{
            if flowerbed[i] == 0 && flowerbed[i+1] == 0{
                avb_plots++
                i++
            }
            continue
        }

        if i == len(flowerbed) -1 {
            if flowerbed[i] == 0 && flowerbed[i-1] == 0{
                avb_plots++
            }   
            continue
        } 

        if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0{
            avb_plots++
            i++
        }
    }
    
    finish:
    return n <= avb_plots

}