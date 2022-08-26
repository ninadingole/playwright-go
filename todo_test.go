package playwright_go_test

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Todo(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err, "Playwright should not error")
	defer func() {
		err := pw.Stop()
		assert.NoError(t, err, "failed to stop playwright")
	}()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{})
	assert.NoError(t, err, "failed to launch browser")
	defer func() {
		err = browser.Close()
		assert.NoError(t, err, "failed to close browser")
	}()

	page, err := browser.NewPage()
	assert.NoError(t, err, "failed to create page")

	_, err = page.Goto("https://todomvc.com/examples/react/#/")
	assert.NoError(t, err, "failed to navigate to page")

	assertTodoItemCountText := func(count int) {
		itemPlural := "items"
		if count == 1 {
			itemPlural = "item"
		}

		todoCounts, err := page.QuerySelector("span.todo-count")
		assert.NoError(t, err, "failed to get todo count")
		assert.NotNil(t, todoCounts, "todo counts not found")

		textContent, err := todoCounts.TextContent()
		assert.NoError(t, err, "failed to get text content")
		assert.Equal(t, fmt.Sprintf("%d %s left", count, itemPlural), textContent)
	}

	assert.NoError(t, page.Click("input.new-todo"), "failed to click new todo")
	assert.NoError(t, page.Type("input[placeholder='What needs to be done?']", "Buy milk"), "failed to type new todo")
	assert.NoError(t, page.Press("input.new-todo", "Enter"), "failed to press enter")

	selector, err := page.QuerySelector("text='Buy milk'")
	assert.NoError(t, err, "failed to get selector")

	content, err := selector.TextContent()
	assert.NotNil(t, content, "text content not found")

	assertTodoItemCountText(1)

	activeTab, err := page.QuerySelector("text='Active'")
	assert.NoError(t, err, "failed to get active tab")

	assert.NoError(t, activeTab.Click())

	assertTodoCount := func(count int) {
		todoListCount, err := page.EvalOnSelectorAll("ul.todo-list > li", "el => el.length")
		assert.NoError(t, err, "failed to get todo list count")
		assert.Equal(t, count, todoListCount)
	}

	assertTodoCount(1)

	checkbox, err := page.QuerySelector("input.toggle")
	assert.NoError(t, err, "failed to get checkbox")

	assert.NoError(t, checkbox.Click())

	completedTab, err := page.QuerySelector("text='Completed'")
	assert.NoError(t, err, "failed to get completed tab")

	assert.NoError(t, completedTab.Click())

	assertTodoCount(1)

	clearCompletedBtn, err := page.QuerySelector("text='Clear completed'")
	assert.NoError(t, err, "failed to get clear completed button")

	assert.NoError(t, clearCompletedBtn.Click())

	assertTodoCount(0)
}
