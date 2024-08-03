package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("\n Name: %v\n", ctx.Value("name"))
	anotherCtx := context.WithValue(ctx, "name", "Nishu") // It doesn't modify existing context instead it wraps the existing context.
	doSomethingElse(anotherCtx)

}
func doSomethingElse(ctx context.Context) {
	fmt.Printf("\n Name: %v\n", ctx.Value("name"))

}

func main() {
	//ctx := context.TODO() // This is how we start an empty context.
	ctx := context.Background() // This is another way how we start an empty context.
	ctx = context.WithValue(ctx, "name", "Kapil")
	doSomething(ctx)

}
