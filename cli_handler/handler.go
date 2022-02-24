package cli_handler

import (
	"fmt"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
)

const (
	ActionProcessMoney = "process_money"
	ActionProcessItems = "process_items"
)

type (
	cliHandler struct {
		storage storage.Storage
	}
)

func (h *cliHandler) ProcessMoneyRewards() error {
	rewards, err := h.storage.GetUnprocessedMoneyRewards()
	if err != nil {
		return err
	}

	var ids = make([]int64, len(rewards))
	for i, reward := range rewards {
		fmt.Printf(
			"Processed money reward %d:%d:%d\n",
			reward.GetId(), reward.GetUserId(), reward.GetAmount(),
		)
		ids[i] = reward.GetId()
	}
	return h.storage.SetMoneyRewardsProcessed(ids)
}

func (h *cliHandler) ProcessItemRewards() error {
	rewards, err := h.storage.GetUnprocessedItemsRewards()
	if err != nil {
		return err
	}

	var ids = make([]int64, len(rewards))
	for i, reward := range rewards {
		fmt.Printf(
			"Processed item reward %d:%d:%s\n",
			reward.GetId(), reward.GetUserId(), reward.GetType(),
		)
		ids[i] = reward.GetId()
	}
	return h.storage.SetItemsRewardsProcessed(ids)
}

func MakeCliHandler(storage storage.Storage) *cliHandler {
	return &cliHandler{storage: storage}
}
