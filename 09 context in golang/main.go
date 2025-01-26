package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("name"))
}

func main() {
	//The context.Background function creates an empty context like context."TODO" does, but it’s designed to be used where you intend to start a known context. Fundamentally the two functions do the same thing: they return an empty context that can be used as a context.Context. The biggest difference is how you signal your intent to other developers. If you’re unsure which one to use, context.Background is a good default option.
	ctx := context.Background() // We can use this when we don't know which context to use.
	ctx = context.WithValue(ctx, "name", "kapil")
	doSomething(ctx)
}
