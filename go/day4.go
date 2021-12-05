package aoc

import (
	"strconv"
	"strings"
)

type bingoCard = []map[int]bool

func RunDay4(path string) (int, int) {
	lines := getLines(path, "inputs/day4.txt")
	numbers := parseBingoNumbers(lines[0])
	cards := parseBingoCards(lines)

	return getScoreForFinishedCard(cards, numbers, 1), getScoreForFinishedCard(cards, numbers, len(cards))
}

func getScoreForFinishedCard(cards []bingoCard, numbers []int, gamesToPlay int) int {
	finished := make(map[int]bool)
	for _, number := range numbers {
		for i, card := range cards {
			done := false
			for _, row := range card {
				delete(row, number)
				if len(row) == 0 {
					done = true
				}
			}
			if done {
				finished[i] = true
			}
			if len(finished) == gamesToPlay {
				return calculateBingoScore(card, number)
			}
		}
	}

	return -1
}

func calculateBingoScore(card bingoCard, number int) int {
	sum := 0
	for _, row := range card {
		for n := range row {
			sum += n
		}
	}

	return sum * number / 2
}

func parseBingoNumbers(line string) []int {
	var numbers []int
	items := strings.Split(line, ",")
	for _, item := range items {
		number, _ := strconv.Atoi(item)
		numbers = append(numbers, number)
	}

	return numbers
}

func parseBingoCards(lines []string) []bingoCard {
	var cards []bingoCard

	numberOfCards := len(lines) / 6

	for i := 0; i < numberOfCards; i++ {
		cardLines := lines[(i*6 + 2):(i*6 + 7)]
		cards = append(cards, parseBingoCard(cardLines))
	}

	return cards
}

func parseBingoCard(lines []string) bingoCard {
	card := make(bingoCard, 10)
	for i := 0; i < 10; i++ {
		card[i] = make(map[int]bool)
	}

	for row, line := range lines {
		items := strings.FieldsFunc(line, func(c rune) bool { return c == ' ' })
		for column, item := range items {
			number, _ := strconv.Atoi(item)
			card[row][number] = true
			card[column+5][number] = true
		}
	}

	return card
}
