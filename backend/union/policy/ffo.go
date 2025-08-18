package policy

import (
	"context"

	"github.com/rclone/rclone/backend/union/upstream"
	"github.com/rclone/rclone/fs"
)

func init() {
	registerPolicy("ffo", &FFO{})
}

// FFO stands for first found by order
// Search category: same as epff.
// Action category: same as epff.
// Create category: Given the order of the candidates, act on the first one found.
type FFO struct {
	FF
}

// Action category policy, governing the modification of files and directories
func (p *FFO) Action(ctx context.Context, upstreams []*upstream.Fs, path string) ([]*upstream.Fs, error) {
	if len(upstreams) == 0 {
		return nil, fs.ErrorObjectNotFound
	}
	upstreams = filterRO(upstreams)
	if len(upstreams) == 0 {
		return nil, fs.ErrorPermissionDenied
	}
	return upstreams[:1], nil
}

// SearchEntries is SEARCH category policy but receiving a set of candidate entries
func (p *FFO) SearchEntries(entries ...upstream.Entry) (upstream.Entry, error) {
	if len(entries) == 0 {
		return nil, fs.ErrorObjectNotFound
	}
	// entry := entries[0]
	// if entry.UpstreamFs().SupportDeleteMark() {
	// 	if strings.HasSuffix(entry.Remote(), fs.DeleteSuffix) {
	// 		return entry, fs.ErrorObjectNotFound
	// 	}
	// }
	return entries[0], nil
}
