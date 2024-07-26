package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hoshigakikisame/kabarin/pkg/providers"
	"github.com/hoshigakikisame/kabarin/pkg/utils"
	"github.com/projectdiscovery/gologger"
)

type runner struct {
	options   *Options
	providers *providers.Providers
}

func New(options *Options, providers *providers.Providers) *runner {
	return &runner{options: options, providers: providers}
}

func (r *runner) Notify() error {

	if utils.HasStdin() {
		gologger.Info().Msg("Starting message notifier")

		if r.options.isBulk {
			dataBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			for chunk := range utils.TextChunkStream(string(dataBytes), r.options.CharLimit) {
				r.providers.SendText(&chunk, &o.Delay)
			}
		} else {
			sc := bufio.NewScanner(os.Stdin)

			for sc.Scan() {
				stream := sc.Text()

				for chunk := range utils.TextChunkStream(stream, r.options.CharLimit) {
					r.providers.SendText(&chunk, &o.Delay)
				}

			}
		}
	}

	if r.options.File != "" {
		gologger.Info().Msg("Starting file notifier")

		var outFileName string

		if r.options.ChunkSize == 0 {
			data, err := os.ReadFile(r.options.File)
			if err != nil {
				return err
			}
			outFileName = filepath.Base(r.options.File)

			r.providers.SendFile(&outFileName, &data, &o.Delay)
		} else {
			iteration := 1
			for chunk := range utils.FileChunkStream(r.options.File, r.options.ChunkSize) {
				outFileName := fmt.Sprintf("%s_pt-%d%s", strings.TrimSuffix(filepath.Base(r.options.File), filepath.Ext(r.options.File)), iteration, filepath.Ext(r.options.File))

				r.providers.SendFile(&outFileName, &chunk, &o.Delay)
				iteration++
			}
		}
	}

	return nil
}
