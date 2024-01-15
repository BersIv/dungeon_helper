package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAccountEditBinding

class AccountEdit : Fragment() {

    companion object {
        fun newInstance() = AccountEdit()
    }

    private lateinit var viewModel: AccountEditViewModel

    private var _binding: FragmentAccountEditBinding? = null
    private val binding get() = _binding!!
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val accountEditViewModel = ViewModelProvider(this)[AccountEditViewModel::class.java]
        _binding = FragmentAccountEditBinding.inflate(inflater, container, false)
        val root: View =  binding.root
        val textView: TextView = binding.textAccountEdit
        accountEditViewModel.text.observe(viewLifecycleOwner)
        {
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

    }

}