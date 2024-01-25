package com.example.dungeon_helper.fragments.main.info

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class InfoChapter1ContentViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "INFO_CONTENT"
    }
    val text: LiveData<String> = _text
}