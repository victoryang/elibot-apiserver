package fs

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
)
type eventfunc func(fsnotify.Event, error)

type FileWatcher struct {
	filename 		string
	watcher 		*fsnotify.Watcher
	handle			eventfunc
}

var ctx context.Context
var cancel context.CancelFunc

func (l *FileWatcher) Watch() {
	go func(){
		for {
			select {
			case e := <-l.watcher.Events:
				go l.handle(e, nil)
			case err := <-l.watcher.Errors:
				go l.handle(fsnotify.Event{}, err)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (l *FileWatcher) Close() {
	cancel()
	l.watcher.Close()
} 

func NewFileWatcher(filename string, handle eventfunc) (*FileWatcher, error) {
	ctx, cancel = context.WithCancel(context.Background())

	var w *fsnotify.Watcher
	var err error
	if w, err = fsnotify.NewWatcher(); err!=nil {
		return nil, fmt.Errorf("Could not new watcher: ", err)
	}

	if err = w.Add(filename); err!=nil {
		return nil, fmt.Errorf("Could not add file to watcher: ", err)
	}

	return &FileWatcher{
		logfile:		filename,
		watcher:		w,
		handle:			handle,
	}, nil
}