package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class MasterLobbyInfoContentViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "PLAY_MASTER_INFO_CONTENT"
    }
    val text: LiveData<String> = _text
}