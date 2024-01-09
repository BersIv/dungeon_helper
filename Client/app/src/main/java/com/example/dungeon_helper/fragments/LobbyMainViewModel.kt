package com.example.dungeon_helper.fragments

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class LobbyMainViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "LOBBY"
    }
    val text: LiveData<String> = _text
}