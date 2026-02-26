package interfaces

import (
	"context"
	"log/slog"
)

/*
В пакете slog существует интерфейс Handler для конфигурации обработки логов. Вам необходимо создать структуру CompositeHandler которая позволяет обрабатывать логи сразу в нескольких Handler. CompositHandler содержит под собой []slog.Handler, при вызове методов у CompositeHandler, соответствующие методы должны вызываться у каждого хендлера из слайса.

Для проверки работоспособности CompositeHandler, необходимо создать Logger с помощью CompositeHandler. При логгировании события через данный логгер, логи должны выводиться в терминал пользователя в текстовом формате и в файл в формате json. Для этого испльзуйте стандарные реализации  реализации интерфейса Handler -  JSONHandler и TextHandler.
*/

type CompositeHandler struct {
	Handlers []slog.Handler
}

func NewCompositeHandler(handlers ...slog.Handler) *CompositeHandler {
	return &CompositeHandler  {Handlers: handlers}
}

func (ch *CompositeHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range ch.Handlers {
		if !handler.Enabled(ctx, level) {
			return false
		}
	}
	return true
}

func (ch *CompositeHandler) Handle(ctx context.Context, record slog.Record) error {
	var err error
	for _, handler := range ch.Handlers {
		if err = handler.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

func (ch *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(ch.Handlers))
	for i, handler := range ch.Handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return &CompositeHandler{
		Handlers: newHandlers,
	}
}

func (ch *CompositeHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(ch.Handlers))
	for i, handler := range ch.Handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return &CompositeHandler{
		Handlers: newHandlers,
	}
}


