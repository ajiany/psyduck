package tools

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				if err, ok := v.(error); ok {
					logrus.WithError(err).Error("panic error")
				} else {
					logrus.Errorf("painc: %+v", v)
				}
			}
		}()

		fn()
	}()
}

func Group(ctx context.Context, fs ...func() error) error {
	return LimitGroup(ctx, 10, fs...)
}

func LimitGroup(ctx context.Context, limit int, fs ...func() error) error {
	if len(fs) == 0 {
		return nil
	}

	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(limit)
	for _, f := range fs {
		fn := func(f func() error) func() error {
			return func() error {
				var err error
				defer func() {
					if v := recover(); v != nil {
						if _, ok := v.(error); ok {
							err = v.(error)
						} else {
							err = fmt.Errorf("panic error:%v", v)
						}
					}
				}()

				err = f()

				return err
			}
		}

		eg.Go(fn(f))
	}
	return eg.Wait()
}
