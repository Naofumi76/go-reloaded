package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func fileSelect(s string) string {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(content)
	return goReloaded(strContent)

}

func goReloaded(s string) string {
	result := ""
	rE := regexp.MustCompile(`(\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))|[.|,|?|!|:|;|']{1,}|([\w\d\(\)']{1,})`)
	tab := rE.FindAllString(s, -1)
	applyChanges(tab)
	applyPonct(tab)
	tab = removeCommand(tab)
	for _, value := range tab {
		if value[0] == ',' || value[0] == '.' || value[0] == '!' || value[0] == '?' || value[0] == ':' || value[0] == ';' {
			result += value
		} else {
			result += " " + value
		}
	}
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "  ", " ")
	return result
}

func applyChanges(tab []string) {
	nb := 1
	for indice, value := range tab {
		if string(value[0]) == "(" {
			if value == "(up)" {
				for nb != 0 {
					tab[indice-nb] = ToUpper(tab[indice-nb])
					nb--
				}
			} else if tab[indice][:4] == "(up," {
				nombre := Atoi(tab[indice][5 : len(tab[indice])-1])
				for nombre != 0 {
					tab[indice-nombre] = ToUpper(tab[indice-nombre])
					nombre--
				}
			} else if value == "(cap)" {
				for nb != 0 {
					tab[indice-nb] = Capitalize(tab[indice-nb])
					nb--
				}
			} else if tab[indice][:5] == "(cap," {
				nombre := Atoi(tab[indice][6 : len(tab[indice])-1])
				for nombre != 0 {
					tab[indice-nombre] = Capitalize(tab[indice-nombre])
					nombre--
				}

			} else if value == "(low)" {
				for nb != 0 {
					tab[indice-nb] = ToLower(tab[indice-nb])
					nb--
				}
			} else if tab[indice][:5] == "(low," {
				nombre := Atoi(tab[indice][6 : len(tab[indice])-1])
				for nombre != 0 {
					tab[indice-nombre] = ToLower(tab[indice-nombre])
					nombre--
				}

			} else if value == "(bin)" {
				tab[indice-1] = strconv.Itoa(AtoiBase(tab[indice-1], "01"))
			} else if value == "(hex)" {
				i, _ := strconv.ParseInt(tab[indice-1], 16, 0)
				var temp int = int(i)
				tab[indice-1] = strconv.Itoa(temp)
			}
		}
		nb = 1
	}
}

func applyPonct(tab []string) {
	quote := false
	for indice := 0; indice < len(tab); indice++ {
		value := tab[indice]
		if len(value) == 1 && (value == "a" || value == "A") {
			if indice+1 < len(tab) && (tab[indice+1][0] == 'a' || tab[indice+1][0] == 'e' || tab[indice+1][0] == 'i' || tab[indice+1][0] == 'o' || tab[indice+1][0] == 'u' || tab[indice+1][0] == 'A' || tab[indice+1][0] == 'E' || tab[indice+1][0] == 'O' || tab[indice+1][0] == 'I' || tab[indice+1][0] == 'U') {
				tab[indice] += "n"
			}
		}

		if value == "'" {
			if quote {
				quote = false

				if indice > 0 {
					tab[indice-1] = tab[indice-1] + value

					tab = append(tab[:indice], tab[indice+1:]...)
					indice--
				}
			} else {
				quote = true

				if indice < len(tab)-1 {
					tab[indice+1] = value + tab[indice+1]

					tab = append(tab[:indice], tab[indice+1:]...)
					indice--
				}
			}
		}
	}
}

func SplitWhiteSpaces(s string) []string {
	parenth := false
	var tab []string
	var chaine string
	for index, value := range s {
		if value == '(' {
			parenth = true
		} else if value == ')' {
			parenth = false
		}
		if index == len(s)-1 {
			chaine += string(value)
			tab = append(tab, chaine)
			chaine = ""
		} else if value == '.' || value == ',' || value == '!' || value == '?' || value == ':' || value == ';' {
			if parenth {
				chaine += string(value)
			} else {
				if chaine != "" {
					tab = append(tab, chaine)
					tab = append(tab, string(value))
					chaine = ""
				} else {
					tab = append(tab, string(value))
				}
			}
		} else if string(value) == "'" {
			if chaine == "" {
				tab = append(tab, string(value))
			} else {
				tab = append(tab, chaine)
				tab = append(tab, string(value))
				chaine = ""
			}

		} else if value != ' ' && value != '\t' && value != '\n' {
			chaine += string(value)
		} else if (value == ' ' || value == '\t' || value == '\n') && chaine != "" && !parenth {
			tab = append(tab, chaine)
			chaine = ""
		} else if (value == ' ' || value == '\t' || value == '\n') && chaine != "" && parenth {
			chaine += string(value)
		}
	}
	return tab
}

func removeCommand(tab []string) []string {
	var newTab []string
	for indice, value := range tab {
		if value[0] == '(' || value[len(value)-1] == ')' {
			//indice < len(value) &&
			continue
		} else if string(value[0]) == "'" && len(value) == 1 {
			continue
		} else {
			newTab = append(newTab, tab[indice])
		}
	}
	return newTab
}

func ToLower(s string) string {
	var chaine string
	for _, value := range s {
		if rune(value) >= 65 && rune(value) <= 90 {
			val := value + ('a' - 'A')
			chaine += string(val)
		} else {
			chaine += string(value)
		}
	}
	return chaine
}

func ToUpper(s string) string {
	var chaine string
	for _, value := range s {
		if rune(value) >= 97 && rune(value) <= 122 {
			val := value + ('A' - 'a')
			chaine += string(val)
		} else {
			chaine += string(value)
		}
	}
	return chaine
}

func Capitalize(s string) string {
	tab := []rune(s)
	debut_mot := true
	for index := 0; index != len(s); index++ {
		if IsAlpha(tab[index]) == true && debut_mot {
			if tab[index] >= 'a' && tab[index] <= 'z' {
				tab[index] = ('A' - 'a') + tab[index]
			}
			debut_mot = false
		} else if tab[index] >= 'A' && tab[index] <= 'Z' {
			tab[index] = 'a' - 'A' + tab[index]
		} else if IsAlpha(tab[index]) == false {
			debut_mot = true
		}
	}
	return string(tab)
}

func IsAlpha(r rune) bool {
	if !(r >= 'A' && r <= 'Z') && !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'z') {
		return false
	}
	return true
}

func AtoiBase(s string, base string) int {

	result := 0
	cpt := len(s) - 1
	index := 0
	for _, val := range s {
		for ind, char := range base {
			if char == val {
				index = ind
			}
		}
		result += index * iterativePower(len(base), cpt)
		cpt--
	}
	return result
}

func iterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	for power != 0 {
		result = result * nb
		power--
	}
	return result
}

func Atoi(chaine string) int {
	nombre := 0
	est_negatif := false
	for index, value := range chaine {
		if index == 0 && (value == '+' || value == '-') {
			if value == '-' {
				est_negatif = true
			}
			continue
		}
		if value < '0' || value > '9' {
			return 0
		}
		nombre = nombre*10 + int(value-'0')
	}
	if est_negatif {
		nombre = nombre * -1
	}
	return nombre
}
