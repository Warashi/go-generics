package monad

import "github.com/Warashi/go-generics/types"

type Monad[T, U, MT, MU any] interface {
	Unitter[U, MU]
	Binder[T, MT, MU]
}

type Unitter[U, MU any] interface {
	Unit(value U) MU
}

type Binder[T, MT, MU any] interface {
	Bind(m MT, f types.Applyer[T, MU]) MU
}

func Map[MU, T, U, MT any, Impl Monad[T, U, MT, MU]](impl Impl, src MT, f types.Applyer[T, U]) MU {
	return impl.Bind(src, types.ApplyerFunc[T, MU](func(value T) MU { return impl.Unit(f.Apply(value)) }))
}

func FlatMap[T, MU, MT any, Impl Binder[T, MT, MU]](impl Impl, src MT, f types.Applyer[T, MU]) MU {
	return impl.Bind(src, f)
}
