package context

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type RemoteBranchesContext struct {
	*BasicViewModel[*models.RemoteBranch]
	*ListContextTrait
	*DynamicTitleBuilder
}

var (
	_ types.IListContext    = (*RemoteBranchesContext)(nil)
	_ types.DiffableContext = (*RemoteBranchesContext)(nil)
)

func NewRemoteBranchesContext(
	c *ContextCommon,
) *RemoteBranchesContext {
	viewModel := NewBasicViewModel(func() []*models.RemoteBranch { return c.Model().RemoteBranches })

	getDisplayStrings := func(startIdx int, length int) [][]string {
		return presentation.GetRemoteBranchListDisplayStrings(c.Model().RemoteBranches, c.Modes().Diffing.Ref)
	}

	return &RemoteBranchesContext{
		BasicViewModel:      viewModel,
		DynamicTitleBuilder: NewDynamicTitleBuilder(c.Tr.RemoteBranchesDynamicTitle),
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				View:       c.Views().RemoteBranches,
				WindowName: "branches",
				Key:        REMOTE_BRANCHES_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
				Transient:  true,
			})),
			list:              viewModel,
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *RemoteBranchesContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}

func (self *RemoteBranchesContext) GetSelectedRef() types.Ref {
	remoteBranch := self.GetSelected()
	if remoteBranch == nil {
		return nil
	}
	return remoteBranch
}

func (self *RemoteBranchesContext) GetDiffTerminals() []string {
	itemId := self.GetSelectedItemId()

	return []string{itemId}
}
