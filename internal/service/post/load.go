package post

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diyorich/post-api/internal/model"
	servErr "github.com/diyorich/post-api/internal/service"
	"os"
	"time"
)

func (s *service) Load(ctx context.Context, filePath string) (err error) {
	const op = "postservice.Load"

	file, err := os.Open(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s: %w", op, servErr.ErrFileNotFound)
		}
	}

	defer func() {
		err := file.Close()
		if err != nil {
			err = fmt.Errorf("%s: %w", op, servErr.ErrFileClose)
		}
	}()

	dec := json.NewDecoder(file)

	if _, err = dec.Token(); err != nil {
		fmt.Printf("Error on reading json: %v", err)
		return err
	}

	var savingErrs []error

	for dec.More() {
		var post model.Post
		if err := dec.Decode(&post); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
		saveErr := s.repository.Add(ctx, post)
		if saveErr != nil {
			savingErrs = append(savingErrs, saveErr)
		}

		cancel()
	}

	if len(savingErrs) != 0 {
		return errors.Join(savingErrs...)
	}

	return nil
}
