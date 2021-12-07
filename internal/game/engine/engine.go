package engine

import (
	"container/list"
	"github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	Order int
	Image *ebiten.Image
}

func (l *Layer) DrawOn(image *ebiten.Image) {
	image.DrawImage(l.Image, &ebiten.DrawImageOptions{})
	l.Image.Clear()
}

func (l *Layer) Draw(image *ebiten.Image, options *ebiten.DrawImageOptions) {
	l.Image.DrawImage(image, options)
}

type Engine struct {
	objects  *list.List
	layers   *list.List
	layerMap map[int]*list.Element
}

func NewEngine() *Engine {
	return &Engine{
		objects:  list.New(),
		layers:   list.New(),
		layerMap: make(map[int]*list.Element),
	}
}

func (e *Engine) AddSingle(object Object) {
	item := e.objects.Front()
	for item != nil {
		if item.Value.(Object).Priority() <= object.Priority() {
			e.objects.InsertBefore(object, item)
			return
		}

		item = item.Next()
	}

	e.objects.PushBack(object)
}

func (e *Engine) Add(objects ...Object) {
	for _, object := range objects {
		e.AddSingle(object)
	}
}

func (e *Engine) newLayer(order, width, height int) {
	layer := &Layer{
		Order: order,
		Image: ebiten.NewImage(width, height),
	}

	item := e.layers.Front()
	for item != nil {
		if item.Value.(*Layer).Order < order {
			e.layerMap[order] = e.layers.InsertBefore(layer, item)
			return
		}

		item = item.Next()
	}

	e.layerMap[order] = e.layers.PushBack(layer)
}

func (e *Engine) Draw(image *ebiten.Image) {
	object := e.objects.Front()

	for object != nil {
		o := object.Value.(Object)

		order := o.Layer()
		if _, layerExists := e.layerMap[order]; !layerExists {
			width, height := image.Size()
			e.newLayer(order, width, height)
		}

		o.DrawOn(e.layerMap[o.Layer()].Value.(*Layer).Image)
		object = object.Next()
	}

	layer := e.layers.Front()
	for layer != nil {
		layer.Value.(*Layer).DrawOn(image)
		layer = layer.Next()
	}
}

func (e *Engine) Update(gameInterface interface{}) error {
	var err error

	var itemsToRemove []*list.Element
	object := e.objects.Front()

	for object != nil {
		var shouldRemove bool
		if shouldRemove, err = object.Value.(Object).Update(gameInterface); err != nil {
			return err
		}

		if shouldRemove {
			itemsToRemove = append(itemsToRemove, object)
		}

		object = object.Next()
	}

	for _, toRemove := range itemsToRemove {
		e.objects.Remove(toRemove)
	}

	return nil
}
