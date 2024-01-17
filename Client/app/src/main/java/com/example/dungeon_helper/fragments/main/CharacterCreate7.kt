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
import com.example.dungeon_helper.databinding.FragmentCharacterCreate7Binding

class CharacterCreate7 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate7()
    }

    private lateinit var viewModel: CharacterCreate7ViewModel
    private var _binding:FragmentCharacterCreate7Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate7ViewModel = ViewModelProvider(this)[CharacterCreate7ViewModel::class.java]
        _binding = FragmentCharacterCreate7Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView:TextView = binding.textCharacterCreate7
        characterCreate7ViewModel.text.observe(viewLifecycleOwner)
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
        val saveBtn = binding.saveCharBtn
        backBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate7_to_characterCreate6)
        }
        saveBtn.setOnClickListener {
            val activity = requireActivity() as MainActivity
            val navView = activity.getNavView()
            val menu = navView.menu
            val menuItem1 = menu.findItem(R.id.navigation_info)
            val menuItem2 = menu.findItem(R.id.navigation_character)
            val menuItem3 = menu.findItem(R.id.navigation_lobby)
            val menuItem4 = menu.findItem(R.id.navigation_account)
            menuItem1.isVisible = true
            menuItem2.isVisible = true
            menuItem3.isVisible = true
            menuItem4.isVisible = true
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate7_to_navigation_character)
        }
    }

}