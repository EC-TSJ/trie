/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package trie

import "ec-tsj/core"

const (
	ALPHABET_SIZE = 26
)

type (
	Item = core.T

	trieNode struct {
		childrens [ALPHABET_SIZE]*trieNode
		data      core.T
		isWordEnd bool
	}

	trie struct {
		root *trieNode
		size int
	}
)

func New() *trie {
	return &trie{
		root: &trieNode{},
	}
}

/**
 * Nos dice el tamaño de la lista
 * @return {int}
 */
func (t *trie) Size() int {
	return t.size
}

/**
 * Añade una clave a la lista
 * @param {string}
 * @param {interface{}}
 */
func (t *trie) Add(word string, data ...Item) {
	t.size++
	current := t.root
	for _, a := range word {
		idx := a - 'a'
		if current.childrens[idx] == nil {
			current.childrens[idx] = &trieNode{}
		}
		current = current.childrens[idx]
	}
	current.data = data
	current.isWordEnd = true
}

func (t *trie) _get(key string) *trieNode {
	current := t.root
	for _, a := range key {
		idx := a - 'a'
		if current.childrens[idx] == nil {
			return nil
		}
		current = current.childrens[idx]
	}

	return current
}

/**
 * Devuelve el valor de una clave si está inserta en la lista
 * @param {string}
 * @return {interface{}}
 */
func (t *trie) Get(key string) Item {
	current := t._get(key)
	if current == nil {
		return false
	}

	return current.data
}

/**
 * Pone/Coloca el valor de una clave si está inserta en la lista
 * @param {string}
 * @param {interface{}}
 */
func (t *trie) Set(key string, data Item) {
	current := t._get(key)
	if current == nil {
		return
	}

	current.data = data
}

/**
 * Devuelve una clave si está inserta en la lista ó false en otro caso
 * @param {string}
 * @return {bool}
 */
func (t *trie) Find(word string) bool {
	current := t._get(word)
	if current == nil {
		return false
	}
	if current.isWordEnd {
		return true
	}
	return false
}

/**
 * Remueve una claves si está inserta en la lista
 * @param {string}
 */
func (t *trie) Remove(word string) {
	current := t._get(word)
	if current == nil {
		return
	}
	current.isWordEnd = false
	t.size--
}

/**
 * Nos dice si existe un prefijo en la lista de claves
 * @param {string}
 * @return {bool}
 */
func (t *trie) Prefix(prefix string) bool {
	current := t._get(prefix)
	if current == nil {
		return false
	}

	return true
}

//!+
/**
 * Devuelve las claves que están insertas en la lista
 * Nos da las claves parciales p.ej. "j" -> "john", "jose" y las totales
 * "" -> "dog", "john", "jose", "reme".
 * @param {string} "" ó "j" ó "k"
 * @return {[]string}
 */
func (t *trie) Keys(key string) []string {
	f := make([]string, 0)
	current := t._get(key)
	if current == nil {
		return f
	}
	if current.isWordEnd {
		f = append(f, key)
	}
	//  Núcleo de la operación
	for i, ab := range current.childrens {
		if ab == nil {
			continue
		}
		skey := key + string(rune(i+'a'))
		if current.isWordEnd {
			f = append(f, skey)
		}
		if ab != nil {
			f = t._keys(ab, skey, f)
		}
	}

	return f
}

// Helper función
func (t *trie) _keys(trie *trieNode, key string, words []string) []string {
	for i, a := range trie.childrens {
		if a == nil {
			continue
		}
		skey := key + string(rune(i+'a'))
		if a.isWordEnd {
			words = append(words, skey)
		}
		if a != nil {
			words = t._keys(a, skey, words)
		}
	}

	return words
}

//!-
