package context

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type RemotesContext struct {
	*BasicViewModel[*models.Remote]
	*ListContextTrait
}

var (
	_ types.IListContext    = (*RemotesContext)(nil)
	_ types.DiffableContext = (*RemotesContext)(nil)
)

func NewRemotesContext(c *ContextCommon) *RemotesContext {
	viewModel := NewBasicViewModel(func() []*models.Remote { return c.Model().Remotes })

	getDisplayStrings := func(startIdx int, length int) [][]string {
		return presentation.GetRemoteListDisplayStrings(c.Model().Remotes, c.Modes().Diffing.Ref)
	}

	return &RemotesContext{
		BasicViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				View:       c.Views().Remotes,
				WindowName: "branches",
				Key:        REMOTES_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			})),
			list:              viewModel,
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *RemotesContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}

func (self *RemotesContext) GetDiffTerminals() []string {
	itemId := self.GetSelectedItemId()

	return []string{itemId}
}
