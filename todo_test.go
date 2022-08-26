package playwright_go_test

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Todo(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		SlowMo:   playwright.Float(300)},
	)
	assert.NoError(t, err)
	defer browser.Close()

	page, err := browser.NewPage()
	assert.NoError(t, err)

	_, err = page.Goto("https://todomvc.com/examples/react/#/")
	assert.NoError(t, err)

	assertTodoItemCount := func(count int) {
		itemPlural := "items"
		if count == 1 {
			itemPlural = "item"
		}

		todoCounts, err := page.QuerySelector("span.todo-count")
		assert.NoError(t, err)
		assert.NotNil(t, todoCounts, "todo counts not found")

		textContent, err := todoCounts.TextContent()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%d %s left", count, itemPlural), textContent)
	}

	assert.NoError(t, page.Click("input.new-todo"))
	assert.NoError(t, page.Type("input[placeholder='What needs to be done?']", "Buy milk"))
	assert.NoError(t, page.Press("input.new-todo", "Enter"))

	selector, err := page.QuerySelector("text='Buy milk'")
	assert.NoError(t, err)

	content, err := selector.TextContent()
	assert.NotNil(t, content)

	assertTodoItemCount(1)

	checkbox, err := page.QuerySelector("input.toggle")
	assert.NoError(t, err)

	assert.NoError(t, checkbox.Click())

	activeTab, err := page.QuerySelector("text='Active'")
	assert.NoError(t, err)

	assert.NoError(t, activeTab.Click())

	todoListCount, err := page.EvalOnSelectorAll("ul.todo-list > li", "el => el.length")
	assert.NoError(t, err)
	assert.Equal(t, 0, todoListCount)

}
