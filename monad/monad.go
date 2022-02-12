package monad

import (
	"github.com/Warashi/go-generics/types"
	"github.com/Warashi/go-generics/zero"
)

type Monad[T, U, MT, MU any] interface {
	Unitter[U, MU]
	Binder[T, MT, MU]
}

type AdditiveMonad[T, U, MT, MU any] interface {
	Monad[T, U, MT, MU]
	Zeroer[MU]
	Pluser[MT]
}

type Unitter[U, MU any] interface {
	Unit(value U) MU
}

type Binder[T, MT, MU any] interface {
	Bind(m MT, f types.Function[T, MU]) MU
}

type Zeroer[T any] interface {
	Zero() T
}

type Pluser[MT any] interface {
	Plus(MT, MT) MT
}

func Map[MU, T, U, MT any, Impl Monad[T, U, MT, MU]](impl Impl, src MT, f types.Function[T, U]) MU {
	return impl.Bind(src, types.Closure[T, MU](func(value T) MU { return impl.Unit(f.Apply(value)) }))
}

func FlatMap[T, MU, MT any, Impl Binder[T, MT, MU]](impl Impl, src MT, f types.Function[T, MU]) MU {
	return impl.Bind(src, f)
}

func Filter[T, MT any, Impl AdditiveMonad[T, T, MT, MT]](impl Impl, src MT, f types.Function[T, bool]) MT {
	return impl.Bind(src, types.Closure[T, MT](func(value T) MT {
		if f.Apply(value) {
			return impl.Unit(value)
		}
		return impl.Zero()
	}))
}

func Do[U, MU, T, MT any, Impl Monad[T, U, MT, MU]](impl Impl, src MT, f types.Consumer[T]) {
	impl.Bind(src, types.Closure[T, MU](func(value T) MU {
		f.Accept(value)
		return zero.New[MU]()
	}))
}
