package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class MasterLobbyHpViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "PLAY_MASTER_HP"
    }
    val text: LiveData<String> = _text
}