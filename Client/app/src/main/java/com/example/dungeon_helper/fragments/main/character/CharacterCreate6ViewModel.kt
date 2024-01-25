package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class CharacterCreate6ViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "Character_create_6"
    }
    val text: LiveData<String> = _text
}