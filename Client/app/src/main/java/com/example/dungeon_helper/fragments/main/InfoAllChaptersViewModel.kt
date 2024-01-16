package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class InfoAllChaptersViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "INFO_CHAPTERS"
    }
    val text: LiveData<String> = _text
}