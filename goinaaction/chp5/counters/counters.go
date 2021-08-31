/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: counters.py
@Time: 2021-08-31 19:43
@Last_update: 2021-08-31 19:43
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package counters

type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}