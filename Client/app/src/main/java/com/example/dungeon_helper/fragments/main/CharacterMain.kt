package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.databinding.FragmentCharacterMainBinding

class CharacterMain : Fragment() {

    companion object {
        fun newInstance() = CharacterMain()
    }
    private lateinit var viewModel: CharacterMainViewModel


    private var _binding: FragmentCharacterMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val characterMainViewModel = ViewModelProvider(this)[CharacterMainViewModel::class.java]

        _binding = FragmentCharacterMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textCharacter
        characterMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

}