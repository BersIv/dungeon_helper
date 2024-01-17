package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate6Binding

class CharacterCreate6 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate6()
    }

    private lateinit var viewModel: CharacterCreate6ViewModel
    private var _binding: FragmentCharacterCreate6Binding? = null
    private val binding get() = _binding!!


    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate6ViewModel = ViewModelProvider(this)[CharacterCreate6ViewModel::class.java]
        _binding = FragmentCharacterCreate6Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView:TextView = binding.textCharacterCreate6
        characterCreate6ViewModel.text.observe(viewLifecycleOwner)
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
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate6_to_characterCreate5)
        }
        val forwBtn = binding.forwBtn
        forwBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate6_to_characterCreate7)
        }
    }
}