package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.R

class AuthRestorePwd2 : Fragment() {

    companion object {
        fun newInstance() = AuthRestorePwd2()
    }

    private lateinit var viewModel: AuthRestorePwd2ViewModel

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.fragment_auth_restore_pwd2, container, false)
    }



}