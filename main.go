package main
//Convert one object to 10 objects

type SomeParentResp struct {
    One  [10]string `json:"one"`
    Two  [10]int    `json:"two"`
    Sub1 ElemRes    `json:"sub1"`
    Sub2 ElemRes    `json:"sub2"`
}

type ElemRes struct {
    One  [10]string `json:"one"`
    Two  [10]int    `json:"two"`
}

//should be splited to

type SomeParent struct {
    One string
    Two int
    Sub1 Elem
    Sub2 Elem
}

type Elem struct {
    One string
    Two int
}


func main() {
}