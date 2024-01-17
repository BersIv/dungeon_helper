package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class LobbyListViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "LobbyList"
    }
    val text: LiveData<String> = _text
}